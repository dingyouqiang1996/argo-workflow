import {GetUserInfoResponse, Info, Version} from '../../../models';
import {baseUrl} from '../base';

import requests from './requests';

let info: Promise<Info>; // we cache this globally rather than in localStorage so it is request once per page refresh

export const InfoService = {
    getInfo() {
        if (info) {
            return info;
        }
        info = requests.get(`api/v1/info`).then(res => {
            const info = res.body as Info;
            info.links = info.links.map(link => {
                // relative UI links don't have a protocol so we're going
                // to prefix it with the UI base URL but URLs that don't
                // parse will be treated as absolute like they are before
                // the relative support
                const protocol = URL.parse(link.url)?.protocol || "absolute";
                if (url.protocol === "") {
                    return {
                        ...link,
                        url: baseUrl() + link.url
                    };
                } else {
                    return link;
                }
            });
            return info;
        });
        return info;
    },

    getVersion() {
        return requests.get(`api/v1/version`).then(res => res.body as Version);
    },

    getUserInfo() {
        return requests.get(`api/v1/userinfo`).then(res => res.body as GetUserInfoResponse);
    },

    collectEvent(name: string) {
        return requests.post(`api/v1/tracking/event`).send({name});
    }
};
