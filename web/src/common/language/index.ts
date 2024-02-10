import * as en from './en.json'

const enObj = JSON.parse(JSON.stringify(en));

// Ideally this will be the entry point for setting translations
const language = (language = 'en') => enObj;

export default language;