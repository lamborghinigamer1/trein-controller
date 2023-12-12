import { Component } from '@angular/core';

@Component({
  selector: 'app-upload-mp3',
  standalone: true,
  imports: [],
  templateUrl: './upload-mp3.component.html',
  styleUrl: './upload-mp3.component.css'
})
export class UploadMp3Component {
  private url: string = 'http://192.168.0.147:8080/upload';
  public selectedFile: File | null = null;

  onFileSelected(event: any) {
    this.selectedFile = event.target.files[0];
  }

  uploadFile() {
    if (!this.selectedFile) {
      console.error('No file selected.');
      return;
    }

    const url = this.url;
    const formData = new FormData();
    formData.append('fileupload', this.selectedFile);

    fetch(url, {
      method: 'POST',
      body: formData,
    })
      .then(response => {
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
        return response.text();
      })
      .then(result => console.log(result))
      .catch(error => console.error('Error:', error));
  }
}
