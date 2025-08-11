import { Routes } from '@angular/router';
import { Home } from './features/home/home';
import { CreateRole } from './features/roles/modals/create-role/create-role';
import { CreatePermission } from './features/roles/modals/create-permission/create-permission';

export const routes: Routes = [
    { path: '', component: Home },
    { path: 'auth', loadChildren: () => import('./features/auth/auth-module').then((m) => m.AuthModule) },
    // * OUTLET MODAL * \\
    { path: 'create-role', component: CreateRole, outlet: 'modal' },
    { path: 'create-permission', component: CreatePermission, outlet: 'modal' },
];
