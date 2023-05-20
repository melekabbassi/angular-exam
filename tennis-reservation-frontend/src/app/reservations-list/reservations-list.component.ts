import { Component, OnInit } from '@angular/core';
import { Reservation } from '../interfaces/reservations';
import { ReservationService } from '../services/ReservationService';

@Component({
  selector: 'app-reservations-list',
  templateUrl: './reservations-list.component.html',
  styleUrls: ['./reservations-list.component.css']
})
export class ReservationsListComponent implements OnInit {
  reservations: Reservation[] = [];

  constructor(private reservationService: ReservationService) { }

  ngOnInit(): void {
    this.reservationService.reservations.subscribe(
      reservations => {
        this.reservations = reservations;
      }
    );
  }

  updateReservation(reservation: Reservation) {
    // Here is where you would put your logic to update the reservation.
    // This might involve navigating to a different component where the reservation can be edited,
    // or opening a modal dialog where the changes can be made.
  }

  deleteReservation(reservation: Reservation) {
    // Here is where you would put your logic to delete the reservation.
    // This could be as simple as removing it from the list of reservations and saving the updated list to local storage.
    this.reservations = this.reservations.filter(res => res !== reservation);
    this.reservationService.saveAllReservations(this.reservations);
  }
}
