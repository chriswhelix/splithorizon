import "speaker"
import "dependency"

func main() {
  speaker.NewSpeaker().Speak()
  dependency.Speak()
}