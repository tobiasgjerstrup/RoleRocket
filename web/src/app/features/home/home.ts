import { Component } from '@angular/core';
import { Roles } from '../roles/roles';

@Component({
    selector: 'app-home',
    imports: [Roles],
    templateUrl: './home.html',
    styleUrl: './home.scss',
})
export class Home {}
