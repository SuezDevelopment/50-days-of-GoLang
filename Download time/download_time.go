/*
The download time depends on several factors such as the speed of the server, network congestion, etc. 
Upload time is similar to the download time.
*/
// calculate the estimated download time using the average download speed

package download_time

func DownloadTime(file_size int, download_speed int) int {
  
  //convert the file size and internet speed to bits
  file_size_in_bits := file_size * 1024 * 1024 * 1024 * 8
  download_speed_in_bits_per_second := download_speed * 1024 * 1024
  
  return file_size_in_bits/download_speed_in_bits_per_second
}
