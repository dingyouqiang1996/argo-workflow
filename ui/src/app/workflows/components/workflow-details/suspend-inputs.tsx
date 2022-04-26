import {Select} from 'argo-ui';
import * as React from 'react';
import {Parameter} from '../../../../models';

interface SuspendInputProps {
    parameters: Parameter[];
    nodeId: string;
    setParameter: (key: string, value: string) => void;
}

export const SuspendInputs = (props: SuspendInputProps) => {
    const [parameters, setParameters] = React.useState(props.parameters);

    const setParameter = (key: string, value: string) => {
        props.setParameter(key, value);
        setParameters(previous => {
            return previous.map(parameter => {
                if (parameter.name === key) {
                    parameter.value = value;
                }
                return parameter;
            });
        });
    };

    const renderSelectField = (parameter: Parameter) => {
        return (
            <React.Fragment key={parameter.name}>
                <br />
                <label>{parameter.name}</label>
                <Select
                    value={parameter.value || parameter.default}
                    options={parameter.enum.map(value => ({
                        value,
                        title: value
                    }))}
                    onChange={selected => {
                        setParameter(parameter.name, selected.value);
                    }}
                />
            </React.Fragment>
        );
    };

    const renderInputField = (parameter: Parameter) => {
        return (
            <React.Fragment key={parameter.name}>
                <br />
                <label>{parameter.name}</label>
                <input
                    className='argo-field'
                    defaultValue={parameter.value || parameter.default}
                    onChange={event => {
                        setParameter(parameter.name, event.target.value);
                    }}
                />
            </React.Fragment>
        );
    };

    const renderFields = (parameter: Parameter) => {
        if (parameter.enum) {
            return renderSelectField(parameter);
        }
        return renderInputField(parameter);
    };

    return (
        <div>
            <h2>Modify parameters</h2>
            {parameters.map(renderFields)}
            <br />
            <br />
            Are you sure you want to resume node {props.nodeId} ?
        </div>
    );
};
