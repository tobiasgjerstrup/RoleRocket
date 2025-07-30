import { effect, Injectable, signal } from '@angular/core';

export type TokenRaw = { token: string };

type Token = {
    exp: Date;
    name: string;
} | null;

@Injectable({
    providedIn: 'root',
})
export class Auth {
    constructor() {
        effect(() => {
            const current = this.token();
            if (current) {
                localStorage.setItem('auth_token', JSON.stringify(current));
            } else {
                localStorage.removeItem('auth_token');
            }
        });

        const saved = localStorage.getItem('auth_token');
        if (saved) {
            try {
                const parsed = JSON.parse(saved);

                const expiration = new Date(parsed.exp);
                const now = new Date();
                if (expiration < now) {
                    localStorage.removeItem('auth_token');
                } else {
                    this.token.set(parsed);
                }
            } catch {
                localStorage.removeItem('auth_token');
            }
        }
    }

    public token = signal<Token>(null);

    public authWithToken(token: string) {
        const parsedToken = JSON.parse(atob(token.split('.')[1]));
        parsedToken.exp = new Date(parsedToken.exp * 1000);
        this.token.set(parsedToken);
    }

    public logout() {
        this.token.set(null);
    }
}
