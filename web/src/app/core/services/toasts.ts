import { Injectable, signal } from '@angular/core';

export interface Toast {
    id: number;
    title: string;
    message: string;
    timestamp: string;
    type?: 'success' | 'error' | 'info';
}

@Injectable({
    providedIn: 'root',
})
export class Toasts {
    private toasts = signal<Toast[]>([]);
    private nextId = 0;

    public getToasts = this.toasts;

    public showToast(title: string, message: string, type?: Toast['type']) {
        const toast: Toast = {
            id: this.nextId++,
            title,
            message,
            timestamp: new Date().toLocaleTimeString(),
            type,
        };
        this.toasts.update((list) => [...list, toast]);

        setTimeout(() => this.dismissToast(toast.id), 5000);
    }

    public dismissToast(id: number) {
        this.toasts.update((list) => list.filter((t) => t.id !== id));
    }
}
