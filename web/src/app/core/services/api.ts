import { HttpClient, HttpResponse } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { environment } from '../../environments/environment';
import { Toasts } from './toasts';

@Injectable({
    providedIn: 'root',
})
export class Api {
    constructor(
        private http: HttpClient,
        private toasts: Toasts,
    ) {}

    public async get<T>(url: string): Promise<T> {
        return new Promise<T>((resolve, reject) => {
            this.http.get<T>(environment.apiHost + url, { observe: 'response' }).subscribe({
                next: (res: HttpResponse<T>) => {
                    if (res.status >= 200 && res.status < 300 && res.body !== null) {
                        resolve(res.body);
                    } else {
                        this.toasts.showToast(`Unexpected response: ${res.status}`, 'error');
                        console.error(`Unexpected response: ${res.status}`);
                        reject(`Response not OK: ${res.status}`);
                    }
                },
                error: (err) => {
                    const msg = err?.message || 'Request failed';
                    console.error(`Error: ${msg}`);
                    this.toasts.showToast('Error during GET', `Error: ${msg}`, 'error');
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
                        this.toasts.showToast(`Unexpected response: ${res.status}`, 'error');
                        console.error(`Unexpected response: ${res.status}`);
                        reject(`Response not OK: ${res.status}`);
                    }
                },
                error: (err) => {
                    const msg = err?.message || 'Request failed';
                    console.error(`Error: ${msg}`);
                    this.toasts.showToast('Error during PUT', `Error: ${msg}`, 'error');
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
                        this.toasts.showToast(`Unexpected response: ${res.status}`, 'error');
                        console.error(`Unexpected response: ${res.status}`);
                        reject(`Response not OK: ${res.status}`);
                    }
                },
                error: (err) => {
                    const msg = err?.message || 'Request failed';
                    console.error(`Error: ${msg}`);
                    this.toasts.showToast('Error during POST', `Error: ${msg}`, 'error');
                    reject(err);
                },
            });
        });
    }
}
