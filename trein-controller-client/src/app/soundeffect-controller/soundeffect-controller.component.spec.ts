import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SoundeffectControllerComponent } from './soundeffect-controller.component';

describe('SoundeffectControllerComponent', () => {
  let component: SoundeffectControllerComponent;
  let fixture: ComponentFixture<SoundeffectControllerComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [SoundeffectControllerComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(SoundeffectControllerComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
