import { Component, signal } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { Navbar } from './core/components/navbar/navbar';
import { Toast } from './core/components/toast/toast';

@Component({
    selector: 'app-root',
    imports: [RouterOutlet, Navbar, Toast],
    templateUrl: './app.html',
    styleUrl: './app.scss',
})
export class App {
    protected readonly title = signal('RoleRocket');
}
