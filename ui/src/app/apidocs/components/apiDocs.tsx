import {Page} from 'argo-ui';
import * as React from 'react';
import {useEffect} from 'react';
import SwaggerUI from 'swagger-ui-react';
import 'swagger-ui-react/swagger-ui.css';
import {uiUrl} from '../../shared/base';
import {services} from '../../shared/services';

export const ApiDocs = () => {
    useEffect(() => {
        services.info.collectEvent('openedApiDocs').then();
    }, []);
    return (
        <Page title='Swagger'>
            <div className='argo-container'>
                <SwaggerUI url={uiUrl('assets/openapi-spec/swagger.json')} defaultModelExpandDepth={0} />
            </div>
        </Page>
    );
};
