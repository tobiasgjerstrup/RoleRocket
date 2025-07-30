import { CommonModule } from '@angular/common';
import { Component } from '@angular/core';
import {
    AbstractControl,
    AbstractControlOptions,
    FormBuilder,
    FormGroup,
    ReactiveFormsModule,
    ValidationErrors,
    Validators,
} from '@angular/forms';
import { Auth, TokenRaw } from '../../../core/services/auth';
import { Api } from '../../../core/services/api';

@Component({
    selector: 'app-register',
    imports: [ReactiveFormsModule, CommonModule],
    templateUrl: './register.html',
    styleUrl: './register.scss',
})
export class Register {
    public form: FormGroup;

    ngAfterViewInit() {
        setTimeout(() => {
            this.form.get('password')?.valueChanges.subscribe(() => {
                this.form.get('confirmPassword')?.updateValueAndValidity();
            });
        });
    }
    constructor(
        private fb: FormBuilder,
        private api: Api,
        private auth: Auth,
    ) {
        this.form = this.fb.group(
            {
                username: ['', [Validators.required, Validators.minLength(6), Validators.maxLength(100)]],
                password: ['', [Validators.required, Validators.minLength(6), Validators.maxLength(100)]],
                confirmPassword: ['', [Validators.required, this.confirmPasswordValidator]],
            },
            {
                validators: this.passwordMatchValidator,
            } as AbstractControlOptions,
        );
    }

    public async submit() {
        if (this.form.valid) {
            try {
                await this.api.post('/users', {
                    username: this.form.value.username,
                    password: this.form.value.password,
                });
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

    private passwordMatchValidator(group: FormGroup): ValidationErrors | null {
        const pass = group.get('password')?.value;
        const confirm = group.get('confirmPassword')?.value;
        return pass === confirm ? null : { passwordMismatch: true };
    }

    private confirmPasswordValidator = (control: AbstractControl): ValidationErrors | null => {
        if (!this.form) return null; // avoid errors before form is initialized
        const password = this.form.get('password')?.value;
        const confirm = control.value;
        return password === confirm ? null : { passwordMismatch: true };
    };
}
