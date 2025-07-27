import { Component } from '@angular/core';
import { Auth } from '../auth';
import { RouterModule } from '@angular/router';
import { environment } from '../../environments/environment';

@Component({
    selector: 'app-navbar',
    imports: [RouterModule],
    templateUrl: './navbar.html',
    styleUrl: './navbar.scss',
})
export class Navbar {
    constructor(public auth: Auth) {}

    public readonly version = environment.version;
}
