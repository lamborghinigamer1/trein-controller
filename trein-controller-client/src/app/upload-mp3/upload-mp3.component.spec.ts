import { ComponentFixture, TestBed } from '@angular/core/testing';

import { UploadMp3Component } from './upload-mp3.component';

describe('UploadMp3Component', () => {
  let component: UploadMp3Component;
  let fixture: ComponentFixture<UploadMp3Component>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [UploadMp3Component]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(UploadMp3Component);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
