import { Component } from '@angular/core';
import { Router } from '@angular/router';

@Component({
    selector: 'app-create-role',
    imports: [],
    templateUrl: './create-role.html',
    styleUrl: './create-role.scss',
})
export class CreateRole {
    constructor(private router: Router) {}

    closeModal() {
        this.router.navigate([{ outlets: { modal: null } }]);
    }
}
