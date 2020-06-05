import * as React from 'react';

import * as models from '../../../../models';
import {services} from '../../../shared/services';
import {FullHeightLogsViewer} from './full-height-logs-viewer';
import {Checkbox} from "argo-ui/src/index";

require('./workflow-logs-viewer.scss');

interface WorkflowLogsViewerProps {
    workflow: models.Workflow;
    nodeId: string;
    container: string;
    archived: boolean;
}

interface WorkflowLogsViewerState {
    error?: Error;
    lines: string[];
    followLogStream: boolean;
}

export class WorkflowLogsViewer extends React.Component<WorkflowLogsViewerProps, WorkflowLogsViewerState> {
    constructor(props: WorkflowLogsViewerProps) {
        super(props);
        this.state = {
            lines: [],
            followLogStream: true
        };
    }

    public componentDidMount(): void {
        services.workflows.getContainerLogs(this.props.workflow, this.props.nodeId, this.props.container, this.props.archived, this.state.followLogStream).subscribe(
            log => {
                if (log) {
                    this.setState(state => {
                        log.split('\n').forEach(line => {
                            state.lines.push(line);
                        });
                        return state;
                    });
                }
            },
            error => {
                this.setState({error});
            }
        );
    }

    public render() {
        return (
            <div className='workflow-logs-viewer'>
                <Checkbox
                    checked={this.state.followLogStream}
                    id='followLogStream'
                    onChange={() => {
                        this.setState({followLogStream: !this.state.followLogStream});
                    }}
                />{' '}
                <label htmlFor='followLogStream'>Stream Logs</label>
                <h3>Logs</h3>
                {this.props.archived && (
                    <p>
                        <i className='fa fa-exclamation-triangle' /> Logs for archived workflows maybe overwritten by a more recent workflow with the same name.
                    </p>
                )}
                <p>
                    <i className='fa fa-box' /> {this.props.nodeId}/{this.props.container}
                    {this.state.lines.length > 0 && <small className='muted'> {this.state.lines.length} line(s)</small>}
                </p>
                <div className='white-box'>
                    {this.state.error && (
                        <p>
                            <i className='fa fa-exclamation-triangle status-icon--failed' /> Failed to load logs: {this.state.error.message}
                        </p>
                    )}
                    {!this.state.error && this.state.lines.length === 0 && this.isCurrentNodeRunningOrPending() && (
                        <p>
                            <i className='fa fa-circle-notch fa-spin' /> Waiting for data...
                        </p>
                    )}
                    {!this.state.error && this.state.lines.length === 0 && !this.isCurrentNodeRunningOrPending() && <p>Pod did not output any logs.</p>}
                    {this.state.lines.length > 0 && (
                        <div className='log-box'>
                            <FullHeightLogsViewer
                                source={{
                                    key: `${this.props.workflow.metadata.name}-${this.props.container}`,
                                    loadLogs: () => {
                                        return services.workflows
                                            .getContainerLogs(this.props.workflow, this.props.nodeId, this.props.container, this.props.archived, this.state.followLogStream)
                                            .map(log => {
                                                return log ? log + '\n' : '';
                                            });
                                    },
                                    shouldRepeat: () => this.isCurrentNodeRunningOrPending()
                                }}
                            />
                        </div>
                    )}
                </div>
                {this.state.lines.length === 0 && (
                    <p>
                        Still waiting for data or an error? Try getting{' '}
                        <a href={services.workflows.getArtifactLogsUrl(this.props.workflow, this.props.nodeId, this.props.container, this.props.archived)}>
                            logs from the artifacts
                        </a>
                        .
                    </p>
                )}
            </div>
        );
    }

    private isCurrentNodeRunningOrPending(): boolean {
        return this.props.workflow.status.nodes[this.props.nodeId].phase === 'Running' || this.props.workflow.status.nodes[this.props.nodeId].phase === 'Pending';
    }
}
