import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GroundsComponent } from './grounds.component';

describe('GroundsComponent', () => {
  let component: GroundsComponent;
  let fixture: ComponentFixture<GroundsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ GroundsComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(GroundsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
