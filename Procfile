controller: RETENTION_GC_PERIOD=10s PNS_PRIVILEGED=true DEFAULT_REQUEUE_TIME=${DEFAULT_REQUEUE_TIME} LEADER_ELECTION_IDENTITY=local ALWAYS_OFFLOAD_NODE_STATUS=${ALWAYS_OFFLOAD_NODE_STATUS} OFFLOAD_NODE_STATUS_TTL=30s WORKFLOW_GC_PERIOD=30s UPPERIO_DB_DEBUG=${UPPERIO_DB_DEBUG} ARCHIVED_WORKFLOW_GC_PERIOD=30s ./dist/workflow-controller --executor-image ${IMAGE_NAMESPACE}/argoexec:${VERSION} --namespaced=${NAMESPACED} --namespace ${NAMESPACE} --loglevel ${LOG_LEVEL}
argo-server: UPPERIO_DB_DEBUG=${UPPERIO_DB_DEBUG} ./dist/argo --loglevel ${LOG_LEVEL} server --namespaced=${NAMESPACED} --namespace ${NAMESPACE} --auth-mode ${AUTH_MODE} --secure=$SECURE --x-frame-options=SAMEORIGIN
ui: yarn --cwd ui install && yarn --cwd ui start
