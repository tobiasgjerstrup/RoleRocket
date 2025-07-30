import packageJson from '../../../package.json';

export const environment = {
    version: packageJson.version,
    apiHost: 'http://localhost:8080',
};
