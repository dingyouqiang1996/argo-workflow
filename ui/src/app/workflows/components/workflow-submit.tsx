import {Formik} from 'formik';
import * as jsYaml from 'js-yaml';
import * as React from 'react';
import * as models from '../../../models';
import {AppContext} from '../../shared/context';
import {services} from '../../shared/services';

interface WorkflowSubmitProps {
    defaultWorkflow: models.Workflow;
}

interface WorkflowSubmitState {
    wf: models.Workflow;
    wfString: string;
    error?: Error;
}

export class WorkflowSubmit extends React.Component<WorkflowSubmitProps, WorkflowSubmitState> {
    constructor(props: WorkflowSubmitProps) {
        super(props);
        this.state = {
            wf: this.props.defaultWorkflow,
            wfString: jsYaml.dump(this.props.defaultWorkflow)
        };
    }

    public render() {
        return (
            <div>
                <Formik
                    initialValues={{wf: this.state.wf, wfString: this.state.wfString}}
                    onSubmit={(values, {setSubmitting}) => {
                        services.workflows
                            .create(values.wf, values.wf.metadata.namespace)
                            .then(wf => this.appContext.apis.navigation.goto(`/workflows/${wf.metadata.namespace}/${wf.metadata.name}`))
                            .catch(error => this.setState({error}));
                        setTimeout(() => {
                            alert(JSON.stringify(values, null, 2));
                            setSubmitting(false);
                        }, 400);
                    }}>
                    {(formikApi: any) => (
                        <form onSubmit={formikApi.handleSubmit}>
                            <div className='white-box editable-panel'>
                                <h4>Submit New Workflow</h4>
                                <button type='submit' className='argo-button argo-button--base' disabled={formikApi.isSubmitting}>
                                    Submit
                                </button>
                                {this.state.error && (
                                    <p>
                                        <i className='fa fa-exclamation-triangle status-icon--failed' /> {this.state.error.message}
                                    </p>
                                )}
                                <textarea
                                    name={'wfString'}
                                    className='yaml'
                                    value={formikApi.values.wfString}
                                    onChange={e => {
                                        formikApi.handleChange(e);
                                    }}
                                    onBlur={e => {
                                        formikApi.handleBlur(e);
                                        try {
                                            formikApi.setFieldValue('wf', jsYaml.load(e.currentTarget.value));
                                        } catch (e) {
                                            // Do nothing, validation when Workflow is submitted
                                        }
                                    }}
                                    onFocus={e => (e.currentTarget.style.height = e.currentTarget.scrollHeight + 'px')}
                                    autoFocus={true}
                                />

                                {/* Workflow-level parameters*/}
                                {formikApi.values.wf &&
                                    formikApi.values.wf.spec &&
                                    formikApi.values.wf.spec.arguments &&
                                    formikApi.values.wf.spec.arguments.parameters &&
                                    this.renderParameterFields('Workflow Parameters', 'wf.spec.arguments', formikApi.values.wf.spec.arguments.parameters, formikApi)}
                            </div>
                        </form>
                    )}
                </Formik>
            </div>
        );
    }

    private renderParameterFields(sectionTitle: string, path: string, parameters: models.Parameter[], formikApi: any): JSX.Element {
        return (
            <div className='white-box__details' style={{paddingTop: '50px'}}>
                <h5>{sectionTitle}</h5>
                {parameters.map((param: models.Parameter, index: number) => {
                    if (param != null) {
                        return (
                            <div className='argo-form-row'>
                                <label className='argo-label-placeholder' htmlFor={path + '.parameters[' + index + '].value'}>
                                    {param.name}
                                </label>
                                <input
                                    className='argo-field'
                                    key={path + '.parameters[' + index + '].value'}
                                    name={path + '.parameters[' + index + '].value'}
                                    type={'text'}
                                    value={param.value}
                                    onChange={formikApi.handleChange}
                                    onBlur={e => {
                                        formikApi.handleBlur(e);
                                        formikApi.setFieldValue('wfString', jsYaml.dump(formikApi.values.wf));
                                    }}
                                />
                            </div>
                        );
                    }
                })}
            </div>
        );
    }

    private get appContext(): AppContext {
        return this.context as AppContext;
    }
}
