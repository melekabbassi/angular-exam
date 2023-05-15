import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { AuthenticationService } from '../services/authentication.service';
import { Users } from '../interfaces/users'; 

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.css']
})
export class DashboardComponent implements OnInit {
  user!: Users;
  articles = [
    { 
      title: 'Article 1', 
      content: 'This is the content for Article 1', 
      author: 'Author 1',
      image: 'https://c0.wallpaperflare.com/preview/218/795/760/sport-ball-clay-court.jpg' 
    },
    { 
      title: 'Article 2', 
      content: 'This is the content for Article 2', 
      author: 'Author 2',
      image: 'https://c0.wallpaperflare.com/preview/712/777/206/montecarlo-nadal.jpg' 
    },
    // Add more articles here
  ];

  constructor(
    private http: HttpClient,
    private authService: AuthenticationService
  ) {}

  ngOnInit() {
    this.user = this.authService.getCurrentUser();
  }
}
