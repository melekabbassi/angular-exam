import { Component } from '@angular/core';
import { Users } from '../interfaces/users';
import { UserapiService } from '../services/userapi.service';

@Component({
  selector: 'app-validation',
  templateUrl: './validation.component.html',
  styleUrls: ['./validation.component.css']
})
export class ValidationComponent {

  array: Users[] = new Array<Users>();

  constructor(private UserapiService: UserapiService) { }

  deleteUser(id: number) {
    this.UserapiService.delete(id, this.array[id]).subscribe(
      response => {
        console.log(response);
        this.array.splice(id, 1);
      }
    )
  }

  ngOnInit(): void {
    this.UserapiService.getData().subscribe(
      response => {
        this.array = response;
      }
    )
  }

  editUser(id: number) {
    this.UserapiService.update(id, this.array[id]).subscribe(
      response => {
        console.log(response);
      }
    )
  }
}
