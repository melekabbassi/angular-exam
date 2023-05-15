import { Injectable } from '@angular/core';
import { Users }  from '../interfaces/users';

@Injectable({
  providedIn: 'root'
})
export class AuthenticationService {
  private currentUser!: Users;

  setCurrentUser(user: Users) {
    this.currentUser = user;
  }

  getCurrentUser(): Users {
    return this.currentUser;
  }
}
