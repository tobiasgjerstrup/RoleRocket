import { Component, computed } from '@angular/core';
import { Toasts } from '../../services/toasts';
declare var bootstrap: any;

@Component({
    selector: 'app-toast',
    imports: [],
    templateUrl: './toast.html',
    styleUrl: './toast.scss',
})
export class Toast {
    toasts = computed(() => this.toastService.getToasts());

    constructor(private toastService: Toasts) {}
    dismiss(id: number) {
        this.toastService.dismissToast(id);
    }
}
