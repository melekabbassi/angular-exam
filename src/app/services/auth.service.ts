import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Users } from '../interfaces/users';

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  private apiurl = 'http://localhost:42400/users';

  constructor(private http: HttpClient) { }

  //GET all users
  getAllUsers() {
    return this.http.get<Users[]>(this.apiurl)
  }

  // Delete user
  deleteUser(id: number) {
    return this.http.delete(`${this.apiurl}/${id}`);
  }

  // Create user
  register(user: Users): Observable<Users> {
    return this.http.post<Users>(this.apiurl, user);
  }

  // Update user
  update(id: number, user: Users): Observable<Users> {
    return this.http.put<Users>(`${this.apiurl}/${id}`, user);
  }

  // Get user by id
  getById(id: number){
    return this.http.get<Users>(`${this.apiurl}/${id}`);
  }
}
