import { Component } from '@angular/core';
declare var bootstrap: any;

@Component({
    selector: 'app-toast',
    imports: [],
    templateUrl: './toast.html',
    styleUrl: './toast.scss',
})
export class Toast {
    ngAfterViewInit() {
        document.querySelectorAll('.toast').forEach((toastEl) => {
            const toast = new bootstrap.Toast(toastEl);
            toast.show();
        });
    }
}
