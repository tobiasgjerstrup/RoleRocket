import { CommonModule } from '@angular/common';
import { Component } from '@angular/core';
import { FormGroup, FormBuilder, Validators, ReactiveFormsModule, AbstractControlOptions } from '@angular/forms';
import { Auth, TokenRaw } from '../../../core/services/auth';
import { Api } from '../../../core/services/api';

@Component({
    selector: 'app-login',
    imports: [ReactiveFormsModule, CommonModule],
    templateUrl: './login.html',
    styleUrl: './login.scss',
})
export class Login {
    public form: FormGroup;

    constructor(
        private fb: FormBuilder,
        private auth: Auth,
        private api: Api,
    ) {
        this.form = this.fb.group({
            username: ['', [Validators.required, Validators.minLength(6), Validators.maxLength(100)]],
            password: ['', [Validators.required, Validators.minLength(6), Validators.maxLength(100)]],
        } as AbstractControlOptions);
    }

    public async submit() {
        if (this.form.valid) {
            try {
                const res = await this.api.post<TokenRaw>('/users/token', {
                    username: this.form.value.username,
                    password: this.form.value.password,
                });
                this.auth.authWithToken(res.token);
            } catch (err) {
                console.error(err);
            }
        }
    }
}
