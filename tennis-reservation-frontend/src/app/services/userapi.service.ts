import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Users } from '../interfaces/users';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class UserapiService {

  constructor(private http: HttpClient) { }

  apiurl="http://localhost:42400/users";
  
  getData(): Observable<any> {
    return this.http.get<Users[]>(this.apiurl);
  }

  postData(user: any): Observable<any> {
    return this.http.post<Users[]>(this.apiurl, user);
  }

  update(id: any, user: any): Observable<any> {
    return this.http.put<Users[]>(`${this.apiurl}/${id}`, user);
  }

  delete(id: any, user: any): Observable<any> {
    return this.http.delete<Users[]>(`${this.apiurl}/${id}`, user);
  }
  
}
