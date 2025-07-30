import { HttpClient, HttpResponse } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { environment } from '../../environments/environment';

@Injectable({
    providedIn: 'root',
})
export class Api {
    constructor(private http: HttpClient) {}

    public async get<T>(url: string): Promise<T> {
        return new Promise<T>((resolve, reject) => {
            this.http.get<T>(environment.apiHost + url, { observe: 'response' }).subscribe({
                next: (res: HttpResponse<T>) => {
                    if (res.status >= 200 && res.status < 300 && res.body !== null) {
                        resolve(res.body);
                    } else {
                        // this.toastService.show(`Unexpected response: ${res.status}`);
                        console.log(`Unexpected response: ${res.status}`);
                        reject(`Response not OK: ${res.status}`);
                    }
                },
                error: (err) => {
                    const msg = err?.message || 'Request failed';
                    console.log(`Error: ${msg}`);
                    // this.toastService.show(`Error: ${msg}`);
                    reject(err);
                },
            });
        });
    }

    public async put<T>(url: string, body: any): Promise<T> {
        return new Promise<T>((resolve, reject) => {
            this.http.put<T>(environment.apiHost + url, body, { observe: 'response' }).subscribe({
                next: (res: HttpResponse<T>) => {
                    if (res.status >= 200 && res.status < 300 && res.body !== null) {
                        resolve(res.body);
                    } else {
                        // this.toastService.show(`Unexpected response: ${res.status}`);
                        console.log(`Unexpected response: ${res.status}`);
                        reject(`Response not OK: ${res.status}`);
                    }
                },
                error: (err) => {
                    const msg = err?.message || 'Request failed';
                    console.log(`Error: ${msg}`);
                    // this.toastService.show(`Error: ${msg}`);
                    reject(err);
                },
            });
        });
    }

    public async post<T>(url: string, body: any): Promise<T> {
        return new Promise<T>((resolve, reject) => {
            this.http.post<T>(environment.apiHost + url, body, { observe: 'response' }).subscribe({
                next: (res: HttpResponse<T>) => {
                    if (res.status >= 200 && res.status < 300 && res.body !== null) {
                        resolve(res.body);
                    } else {
                        // this.toastService.show(`Unexpected response: ${res.status}`);
                        console.log(`Unexpected response: ${res.status}`);
                        reject(`Response not OK: ${res.status}`);
                    }
                },
                error: (err) => {
                    const msg = err?.message || 'Request failed';
                    console.log(`Error: ${msg}`);
                    // this.toastService.show(`Error: ${msg}`);
                    reject(err);
                },
            });
        });
    }
}
