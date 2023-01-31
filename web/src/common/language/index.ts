import * as en from './en.json'

const enObj = JSON.parse(JSON.stringify(en));

const language = (key: string): string => enObj[key];

export default language;