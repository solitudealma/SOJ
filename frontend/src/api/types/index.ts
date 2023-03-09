export interface ResponseInterface<T> {
    code?: number | string;
    data?: T;
    msg?: string;
}