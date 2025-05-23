export enum HttpClientBaseMethod {
    POST = 'POST',
    GET = 'GET',
    PUT = 'PUT'
}

export interface HttpClientBaseHeader {
    'Content-Type': string,
    'Authorization'?: string
}

export interface HttpClientBaseStatusCode {
    statusCode: number
    statusText: HttpStatusCode
}

export enum HttpStatusCode {
    Continue = 100,
    SwitchingProtocols = 101,
    Processing = 102,
    Ok = 200,
    Created = 201,
    Accepted = 202,
    NonAuthoritativeInformation = 203,
    NoContent = 204,
    ResetContent = 205,
    PartialContent = 206,
    MultiStatus = 207,
    AlreadyReported = 208,
    ImUsed = 226,
    MultipleChoices = 300,
    MovedPermanently = 301,
    Found = 302,
    SeeOther = 303,
    NotModified = 304,
    UseProxy = 305,
    SwitchProxy = 306,
    TemporaryRedirect = 307,
    PermanentRedirect = 308,
    BadRequest = 400,
    Unauthorized = 401,
    PaymentRequired = 402,
    Forbidden = 403,
    NotFound = 404,
    MethodNotAllowed = 405,
    NotAcceptable = 406,
    ProxyAuthenticationRequired = 407,
    RequestTimeout = 408,
    Conflict = 409,
    Gone = 410,
    LengthRequired = 411,
    PreconditionFailed = 412,
    PayloadTooLarge = 413,
    UriTooLong = 414,
    UnsupportedMediaType = 415,
    RangeNotSatisfiable = 416,
    ExpectationFailed = 417,
    IAmATeapot = 418,
    MisdirectedRequest = 421,
    UnprocessableEntity = 422,
    Locked = 423,
    FailedDependency = 424,
    UpgradeRequired = 426,
    PreconditionRequired = 428,
    TooManyRequests = 429,
    RequestHeaderFieldsTooLarge = 431,
    UnavailableForLegalReasons = 451,
    InternalServerError = 500,
    NotImplemented = 501,
    BadGateway = 502,
    ServiceUnavailable = 503,
    GatewayTimeout = 504,
    HttpVersionNotSupported = 505,
    VariantAlsoNegotiates = 506,
    InsufficientStorage = 507,
    LoopDetected = 508,
    NotExtended = 510,
    NetworkAuthenticationRequired = 511,
}

export class HttpClientBase {
    private resp: Response
    private respBody: any

    async createRequest(
        url: string,
        method: HttpClientBaseMethod,
        httpHeader?: HttpClientBaseHeader,
        requestData?: any
    ): Promise<this> {
        const requestInit: RequestInit = {
            method
        }

        if (httpHeader) {
            requestInit.headers = httpHeader as any
        }

        if (requestData) {
            requestInit.body = JSON.stringify(requestData)
        }

        this.resp = await fetch(url, requestInit)
        this.respBody = await this.resp.json()

        return this
    }

    statusCode(): HttpClientBaseStatusCode {
        const statusCode = this.resp.status as HttpStatusCode
        const statusText = HttpStatusCode[this.resp.status] as unknown as HttpStatusCode

        return {
            statusCode,
            statusText
        }
    }

    isOk(): boolean {
        return this.resp.ok && !("errors" in this.respBody)
    }

    responseByKey<T>(key: string | null = null): T {
        if (key) {
            return this.respBody[key]
        } else {
            return this.respBody
        }
    }

    responseData<T>(): T {
        return this.respBody.data
    }

    responseErrors(): GQLError | null {
        if ("errors" in this.respBody) {
            return this.respBody.errors;
        }
        return null;
    }
}

interface Location {
    line: string;
    column: string;
}

export interface GQLError {
    message: string;
    locations?: Location[];
    path?: string[];
    extentions?: Map<string, number>;
}

export interface GQLResponse<T> {
    data: T;
    errors?: GQLError;
    isOk: boolean;
}

export class GQLClient extends HttpClientBase {
    private header: HttpClientBaseHeader;

    constructor() {
        super();
        this.header = {
            "Content-Type": "application/json"
        };
    }

    async post<T>(
        url: string,
        requestData: any
    ): Promise<GQLResponse<T>> {
        const request = await this.createRequest(
            url,
            HttpClientBaseMethod.POST,
            this.header,
            requestData
        );

        return {
            data: request.responseData(),
            errors: request.responseErrors() || undefined,
            isOk: request.isOk()
        };
    }
}

export function buildQuery() {


}