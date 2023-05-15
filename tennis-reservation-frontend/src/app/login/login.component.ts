import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { AuthenticationService } from '../services/authentication.service';
import { Users }  from '../interfaces/users';


@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  public loginForm!: FormGroup

  constructor(private formBuilder: FormBuilder, private http: HttpClient, private router: Router, private authService: AuthenticationService) {}

  ngOnInit(): void {
    this.loginForm = this.formBuilder.group({
      email: [''],
      password: ['', Validators.required]
    })
  }

  login(){
    this.http.get<any>('http://localhost:42400/users')
      .subscribe(
        response => {
          const user = response.find((user: any) => {
            return user.email === this.loginForm.value.email && user.password === this.loginForm.value.password;
          });
          if (user) {
            alert('Login successful');
            this.authService.setCurrentUser(user);
            this.loginForm.reset();
            this.router.navigate(['/dashboard']);
          } else {
            alert('Login failed');
          }
    }, error => {
      console.log(error);
    })
  }
}
