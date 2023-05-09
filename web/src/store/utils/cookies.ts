import Cookies, { CookieAttributes } from 'js-cookie';
import jwtDecode, { JwtPayload } from 'jwt-decode';

export const JWT_PROPERTY_NAME: string = 'session'
export const COMMON_COOKIE_CONFIG: CookieAttributes = { expires: 1, sameSite: 'strict' }

export const setUserCookie = (token: string) => {
    const decodedJWT: JwtPayload = jwtDecode(token, { header: true })
    const expiration = decodedJWT?.exp ? new Date(decodedJWT?.exp) : new Date();
    Cookies.set(JWT_PROPERTY_NAME, token, {...COMMON_COOKIE_CONFIG, expires: expiration});
}

export const getUserCookie = () => Cookies.get(JWT_PROPERTY_NAME);

export const removeUserCookie = () => Cookies.remove(JWT_PROPERTY_NAME,  {...COMMON_COOKIE_CONFIG});