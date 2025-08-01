import { Component } from '@angular/core';
import { Api } from '../../core/services/api';

type PermissionsT = Array<{ id: number; name: string; createTime: Date }>;
type RolesT = Array<{ id: number; name: string; createTime: Date; permissions: PermissionsT }>;

@Component({
    selector: 'app-roles',
    imports: [],
    templateUrl: './roles.html',
    styleUrl: './roles.scss',
})
export class Roles {
    constructor(private api: Api) {}

    private roles: RolesT | null = null;

    async ngOnInit() {
        this.roles = await this.api.get<RolesT>('/roles?$embed=permissions');
    }
}
