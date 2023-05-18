import { Injectable } from '@angular/core';
import { Users }  from '../interfaces/users';

@Injectable({
  providedIn: 'root'
})
export class AuthenticationService {
  private currentUser!: Users;

  setCurrentUser(user: Users) {
    localStorage.setItem('user', JSON.stringify(user));
    this.currentUser = user;
  }

  getCurrentUser(): Users {
    const user = localStorage.getItem('user');

    if (user) {
      this.currentUser = JSON.parse(user);
    }

    return this.currentUser;
  }
}
