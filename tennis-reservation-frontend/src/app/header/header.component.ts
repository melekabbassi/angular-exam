import { Component } from '@angular/core';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.css']
})
export class HeaderComponent {
  isLoggedIn(): boolean {
    // Implement your logic to determine the user's login status here
    // You can use local storage, session storage, or any other mechanism to store login information
    // For example, you can check if a token or user information exists in the storage
    // and return true if the user is logged in, or false otherwise
    const token = localStorage.getItem('token');
    return !!token; // Return true if token exists, false otherwise
  }
}
