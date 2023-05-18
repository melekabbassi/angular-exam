import { Component } from '@angular/core';

@Component({
  selector: 'app-services-tennis',
  templateUrl: './services-tennis.component.html',
  styleUrls: ['./services-tennis.component.css']
})
export class ServicesTennisComponent {
  services = [
    {
      title: 'Tennis Court Rentals',
      description: 'Rent our well-maintained tennis courts for your enjoyment and practice.',
      moreInfo: 'Learn about our court availability, rates, and booking process.',
      image: 'https://gamo.ge/wp-content/uploads/2022/04/DSC8694-1-1024x683.jpg',
      isSelected: false
    },
    {
      title: 'Tennis Coaching',
      description: 'Improve your tennis skills with personalized coaching from our experienced instructors.',
      moreInfo: 'Explore our coaching programs, private lessons, and group training options.',
      image: 'https://www.vidatennis.com.au/wp-content/uploads/2020/09/vida-tennis-blogs-25-1024x683.jpg',
      isSelected: false
    },
    {
      title: 'Tennis Equipment Sales',
      description: 'Find the perfect tennis equipment to enhance your game at our pro shop.',
      moreInfo: 'Discover our selection of high-quality racquets, balls, bags, and other tennis essentials.',
      image: 'https://tennisracketball.com/wp-content/uploads/2022/04/different_types_of_tennis_rackets_based_on_material-1024x683.jpg',
      isSelected: false
    },
    {
      title: 'Tennis Tournaments',
      description: 'Participate in thrilling tennis tournaments and test your abilities against other players.',
      moreInfo: 'Stay updated on upcoming tournaments, registration details, and tournament rules.',
      image: 'https://championnatsbanquenationale.com/wp-content/uploads/2022/08/Stade_SJC_841-1024x683.jpg',
      isSelected: false
    },
    {
      title: 'Tennis Social Events',
      description: 'Join our tennis community and enjoy socializing with fellow players in a friendly atmosphere.',
      moreInfo: 'Stay informed about our social events, such as round-robin matches and social mixers.',
      image: 'https://www.vidatennis.com.au/wp-content/uploads/2020/08/vida-tennis-image-17-1024x683.jpg',
      isSelected: false
    },
    {
      title: 'Tennis Pro Shop',
      description: 'Browse through our pro shop to find the latest tennis fashion and branded merchandise.',
      moreInfo: 'Explore our collection of tennis clothing, footwear, accessories, and gift items.',
      image: 'https://www.bsacontractors.com/wp-content/uploads/2019/11/bsa-construction-port-comm-woodfield-tennis-img-7-1024x683.jpg',
      isSelected: false
    },
    {
      title: 'Tennis Stringing Services',
      description: 'Keep your racquet in top condition with our expert stringing services using high-quality strings.',
      moreInfo: 'Learn about our stringing options, turnaround time, and pricing for racquet stringing.',
      image: 'https://tennisracketball.com/wp-content/uploads/2022/10/stringing_a_tennis_strings-1024x683.jpg',
      isSelected: false
    },
    {
      title: 'Tennis Fitness Training',
      description: 'Enhance your on-court performance with our tailored fitness training programs for tennis players.',
      moreInfo: ' Discover our fitness classes, personalized training plans, and professional trainers.',
      image: 'https://www.vidatennis.com.au/wp-content/uploads/2021/08/Social-Tennis-1024x683.png',
      isSelected: false
    },
  ];

  selectCard(index: number) {
    this.services[index].isSelected = !this.services[index].isSelected;
  }

}
