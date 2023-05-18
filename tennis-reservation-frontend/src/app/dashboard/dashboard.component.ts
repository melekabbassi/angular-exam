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
  formData = {email: '', subject: '', message: ''};
  articles = [
    { 
      title: 'The Importance of Footwork in Tennis', 
      content: 'Content: Footwork plays a crucial role in tennis. It determines a player s ability to move efficiently on the court, reach the ball in time, and maintain balance during shots. Good footwork allows players to position themselves properly, enabling them to execute strokes with precision and power. It helps in generating momentum, transferring weight effectively, and recovering quickly after each shot. Footwork drills and exercises are essential for improving agility, speed, and court coverage. Developing strong footwork skills not only enhances a player s overall performance but also helps prevent injuries by reducing stress on the body. ', 
      author: 'John Smith',
      image: 'https://c0.wallpaperflare.com/preview/218/795/760/sport-ball-clay-court.jpg' 
    },
    { 
      title: 'The Mental Game: Mastering Tennis Psychology', 
      content: 'Tennis is not just a physical game; it is also a mental battle. Mental strength and psychological resilience are vital for success on the court. Players need to maintain focus, handle pressure, and stay composed during matches. Developing strategies to manage nerves, deal with distractions, and maintain a positive mindset is crucial. Visualization techniques, breathing exercises, and mental rehearsal can help players stay calm and confident in high-pressure situations. Mental training, alongside physical training, is essential for optimizing performance and unlocking a player s full potential in tennis.', 
      author: 'Emily Johnson',
      image: 'https://c0.wallpaperflare.com/preview/712/777/206/montecarlo-nadal.jpg' 
    },
    
  ];

  constructor(
    private http: HttpClient,
    private authService: AuthenticationService
  ) {}

  ngOnInit() {
    this.user = this.authService.getCurrentUser();
  }

  submitForm() {
    const emailData = {
      to: 'amiraknani@gmail.com',
      from: this.formData.email,
      subject: this.formData.subject,
      text: this.formData.message,
    };
    this.http.post('https://yourserver.com/send-email', emailData).subscribe(response => {
      console.log(response);
    }, error => {
      console.log(error);
    });
  }
}
