import { Component } from '@angular/core';
import { Router } from '@angular/router';

@Component({
    selector: 'app-create-permission',
    imports: [],
    templateUrl: './create-permission.html',
    styleUrl: './create-permission.scss',
})
export class CreatePermission {
    constructor(private router: Router) {}

    closeModal() {
        this.router.navigate([{ outlets: { modal: null } }]);
    }
}
