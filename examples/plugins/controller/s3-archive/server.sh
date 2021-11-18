python3 -m venv /tmp/venv
cd /tmp/venv
source bin/activate

pip3 install boto3

cat <<EOF > server.py
import boto3
import json
from http.server import BaseHTTPRequestHandler, HTTPServer

from botocore.config import Config


class Plugin(BaseHTTPRequestHandler):

    def args(self):
        return json.loads(self.rfile.read(int(self.headers.get('Content-Length'))))

    def reply(self, reply):
        self.send_response(200)
        self.end_headers()
        self.wfile.write(json.dumps(reply).encode("UTF-8"))

    def unsupported(self):
        self.send_response(404)
        self.end_headers()

    def do_POST(self):
        if self.path == '/api/v1/workflow.postOperate':
            args = self.args()
            before = args['old']['metadata'].get('labels', {}).get('workflows.argoproj.io/phase', '')
            after = args['new']['metadata'].get('labels', {}).get('workflows.argoproj.io/phase', '')
            if before == 'Running' and (after == 'Failed' or after == 'Succeeded'):
                s3 = boto3.client('s3',
                                  endpoint_url='http://minio:9000',
                                  aws_access_key_id='admin',
                                  aws_secret_access_key='password',
                                  config=Config(signature_version='s3v4'),
                                  region_name='us-east-1')
                json_bytes = json.dumps(args['new']).encode()
                workflow_name = args['new']['metadata']['name']
                s3.put_object(Body=json_bytes, Bucket='my-bucket', Key=f'workflows/{workflow_name}')
            self.reply({})
        else:
            self.unsupported()


if __name__ == '__main__':
    httpd = HTTPServer(('', 7243), Plugin)
    httpd.serve_forever()
EOF

python3 server.py
