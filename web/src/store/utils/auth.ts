import { CallOptions } from "@bufbuild/connect-web";
import { getUserCookie } from "./cookies";

export const getCredentials = async (): Promise<CallOptions | undefined> => {
    const token = await getUserCookie();
    if (!token) return;

    const options: CallOptions = { headers: { session: token }};
    return options;
};