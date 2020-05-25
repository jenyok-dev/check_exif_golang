package images_exif

import (
  "io"
  "encoding/binary"
)

func getOrientation(r io.ReadSeeker) int {
  if _, err:= r.Seek(12, 0); err != nil {		
    return 0
  }
  var orderTag uint16
  if err := binary.Read(r, binary.BigEndian, &orderTag); err != nil {
    return 0
  }
  var byteOrder binary.ByteOrder
  switch orderTag {
    case 0x4d4d:
      byteOrder = binary.BigEndian
    case 0x4949:
      byteOrder = binary.LittleEndian
    default:
      return 0
  }
  if _, err:= r.Seek(78, 0); err != nil {		
    return 0
  }
  var orientationTag uint16
  if err := binary.Read(r, byteOrder, &orientationTag); err != nil {
    return 0
  }
  if orientationTag > 0 && orientationTag < 9 {
    return int(orientationTag)
  }
  return 0
}
