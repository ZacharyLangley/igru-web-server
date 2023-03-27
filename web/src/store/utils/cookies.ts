import Cookies, { CookieAttributes } from 'js-cookie';

export const JWT_PROPERTY_NAME: string = 'session'
export const COOKIE_CONFIG: CookieAttributes = { expires: 1, sameSite: 'strict' }

export const setUserCookie = (token: string) => Cookies.set(JWT_PROPERTY_NAME, token, COOKIE_CONFIG);

export const getUserCookie = () => Cookies.get(JWT_PROPERTY_NAME);

export const removeUserCookie = () => Cookies.remove(JWT_PROPERTY_NAME, COOKIE_CONFIG);