import { Component, OnInit } from '@angular/core';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { Reservation } from '../interfaces/reservations';
import { ReservationService } from '../services/ReservationService';

@Component({
  selector: 'app-reservation',
  templateUrl: './reservation.component.html',
  styleUrls: ['./reservation.component.css']
})
export class ReservationComponent implements OnInit {
  
  reservationFormGroup: FormGroup = new FormGroup({
    type: new FormControl('', [Validators.required]),
    date: new FormControl('', [Validators.required]),
    hour: new FormControl('', [Validators.required]),
    duration: new FormControl('', [Validators.required]),
  });

  get type() { return this.reservationFormGroup.get('type'); }
  get date() { return this.reservationFormGroup.get('date'); }
  get hour() { return this.reservationFormGroup.get('hour'); }
  get duration() { return this.reservationFormGroup.get('duration'); }

  constructor(private reservationService: ReservationService) { }

  ngOnInit(): void {
  }

  onSubmit() {
    if (this.reservationFormGroup.valid) {
      this.reservationService.saveReservationToLocal(this.reservationFormGroup.value);
      alert('Reservation successful');
      this.reservationFormGroup.reset();
    } else {
      alert('Reservation failed');
    }
  }
}
