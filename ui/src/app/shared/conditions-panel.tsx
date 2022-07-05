import * as React from 'react';
import {Condition, ConditionType} from '../../models';
import {ErrorIcon, WarningIcon} from './components/fa-icons';

interface Props {
    conditions: Condition[];
}

const WarningConditions: ConditionType[] = ['SpecWarning'];
const ErrorConditions: ConditionType[] = ['MetricsError', 'SubmissionError', 'SpecError', 'ArtifactGCError'];

export function hasWarningConditionBadge(conditions: Condition[]): boolean {
    if (conditions.length === 0) {
        return false;
    }

    for (const condition of conditions) {
        if (WarningConditions.includes(condition.type)) {
            return true;
        }
        if (ErrorConditions.includes(condition.type)) {
            return true;
        }
    }

    return false;
}

export function hasArtifactGCError(conditions: Condition[]): boolean {
    if (conditions.length === 0) {
        return false;
    }

    for (const condition of conditions) {
        if (condition.type === 'ArtifactGCError') {
            return true;
        }
    }

    return false;
}

function getConditionIcon(condition: ConditionType): JSX.Element {
    let icon;
    if (WarningConditions.includes(condition as ConditionType)) {
        icon = <WarningIcon />;
    }
    if (ErrorConditions.includes(condition as ConditionType)) {
        icon = <ErrorIcon />;
    }
    if (!icon) {
        return <span />;
    } else {
        return <>{icon}&nbsp;</>;
    }
}

export class ConditionsPanel extends React.Component<Props> {
    public render() {
        return (
            <>
                {this.props.conditions &&
                    Object.entries(this.props.conditions).map(([_, condition]) => {
                        return (
                            <div key={condition.type} style={{lineHeight: '120%', marginTop: '16px'}}>
                                {getConditionIcon(condition.type)}
                                <span className='condition-panel__type'>{condition.type}</span>
                                {': ' + (condition.message || condition.status)}
                            </div>
                        );
                    })}
            </>
        );
    }
}
