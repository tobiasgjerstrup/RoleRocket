import { CommonModule } from '@angular/common';
import { HttpClient } from '@angular/common/http';
import { Component } from '@angular/core';
import { FormGroup, FormBuilder, Validators, ReactiveFormsModule, AbstractControlOptions } from '@angular/forms';
import { Auth, TokenRaw } from '../../../core/auth';

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
        private http: HttpClient,
        private auth: Auth,
    ) {
        this.form = this.fb.group({
            username: ['', [Validators.required, Validators.minLength(5), Validators.maxLength(100)]],
            password: ['', [Validators.required, Validators.minLength(5), Validators.maxLength(100)]],
        } as AbstractControlOptions);
    }

    public submit() {
        if (this.form.valid) {
            this.http
                .post<TokenRaw>('/users/token', {
                    username: this.form.value.username,
                    password: this.form.value.password,
                })
                .subscribe((res) => {
                    this.auth.authWithToken(res.token);
                });
        }
    }
}
