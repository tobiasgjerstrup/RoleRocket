import { Component } from '@angular/core';
import { Api } from '../../core/services/api';
import { Router } from '@angular/router';

type PermissionsT = Array<{ id: number; name: string; createTime: Date }>;
type RolesT = Array<{ id: number; name: string; createTime: Date; permissions: PermissionsT }>;

@Component({
    selector: 'app-roles',
    imports: [],
    templateUrl: './roles.html',
    styleUrl: './roles.scss',
})
export class Roles {
    constructor(
        private api: Api,
        private router: Router,
    ) {}

    public roles: RolesT | null = null;

    async ngOnInit() {
        this.roles = await this.api.get<RolesT>('/roles?$embed=permissions');
    }

    public openCreatePermissionModal() {
        this.router.navigate([{ outlets: { modal: ['create-permission'] } }]);
    }

    public openCreateRoleModal() {
        this.router.navigate([{ outlets: { modal: ['create-role'] } }]);
    }
}
