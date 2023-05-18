import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ServicesTennisComponent } from './services-tennis.component';

describe('ServicesTennisComponent', () => {
  let component: ServicesTennisComponent;
  let fixture: ComponentFixture<ServicesTennisComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ServicesTennisComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ServicesTennisComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
