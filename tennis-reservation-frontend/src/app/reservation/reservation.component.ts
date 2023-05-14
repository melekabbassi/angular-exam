import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Users } from '../interfaces/users';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { AuthService } from '../services/auth.service';

@Component({
  selector: 'app-reservation',
  templateUrl: './reservation.component.html',
  styleUrls: ['./reservation.component.css']
})
export class ReservationComponent implements OnInit {
  array: Users[] = new Array<Users>();

  reservationFormGroup: FormGroup = new FormGroup({
    lastName: new FormControl('', [Validators.required]),
    firstName: new FormControl('', [Validators.required]),
    email: new FormControl('', [Validators.required]),
    password: new FormControl('', [Validators.required, Validators.minLength(6)]),
    profile: new FormControl('', [Validators.required]),
    validated: new FormControl('', [Validators.required])
  });

  get lastName() { return this.reservationFormGroup.get('lastName'); }
  get firstName() { return this.reservationFormGroup.get('firstName'); }
  get email() { return this.reservationFormGroup.get('email'); }
  get password() { return this.reservationFormGroup.get('password'); }
  get profile() { return this.reservationFormGroup.get('profile'); }
  get validated() { return this.reservationFormGroup.get('validated'); }

  addOrPut = false;
  
  constructor(private authService: AuthService, private http: HttpClient) { }

  ngOnInit(): void {
    this.getUser();
  }

  getUser() {
    this.authService.getAllUsers().subscribe(
      (data: Users[]) => {
        this.array = data;
      }
    );
  }

  deleteUser(id: number) {
    this.authService.deleteUser(id).subscribe(
      response => {
        console.log(response);
        this.array = this.array.filter(user => user.id !== id);
      }
    );
  }

  addUser() {
    this.http.post<any>('http://localhost:42400/users', this.reservationFormGroup.value)
      .subscribe(
        response => {
          alert('Registration successful');
          this.reservationFormGroup.reset();
          this.getUser();
        }, error => {
          alert('Registration failed');
          console.log(error);
        })
  }
  
  onSubmit() {
    console.log(this.reservationFormGroup.value);
  }
}
