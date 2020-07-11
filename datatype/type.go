package datatype

import (
	"strconv"

	"github.com/mdouchement/zoia/binary"
)

// Type defines a module type.
type Type uint8

func (t Type) String() string {
	switch t {
	// Interface modules
	case binary.ModuleIDAudioInput:
		return "audio input"
	case binary.ModuleIDAudioOutput:
		return "audio output"
	case binary.ModuleIDPushbutton:
		return "pushbutton"
	case binary.ModuleIDKeyboard:
		return "keyboard"
	case binary.ModuleIDMidiNotesIn:
		return "midi notes in"
	case binary.ModuleIDMidiCCIn:
		return "midi cc in"
	case binary.ModuleIDMidiPressure:
		return "midi pressure"
	case binary.ModuleIDStompswitch:
		return "stompswitch"
	case binary.ModuleIDCportExpCVIn:
		return "cport exp/cv in"
	case binary.ModuleIDCportExpCVOut:
		return "cport exp/cv out"
	case binary.ModuleIDUIButton:
		return "ui button"
	case binary.ModuleIDMidiNoteOut:
		return "midi note out"
	case binary.ModuleIDMidiCCOut:
		return "midi cc out"
	case binary.ModuleIDMidiPCOut:
		return "midi pc out"
	case binary.ModuleIDPixel:
		return "pixel"
	case binary.ModuleIDMidiClockIn:
		return "midi clock in"
	// Audio modules
	case binary.ModuleIDSVFilter:
		return "sv filter"
	case binary.ModuleIDSVAliaser:
		return "aliaser"
	case binary.ModuleIDVCA:
		return "vca"
	case binary.ModuleIDAudionMultiply:
		return "multiply"
	case binary.ModuleIDBitCrusher:
		return "bit crusher"
	case binary.ModuleIDDelayLine:
		return "delay line"
	case binary.ModuleIDOscillator:
		return "oscillator"
	case binary.ModuleIDMultiFilter:
		return "multi-filter"
	case binary.ModuleIDBufferDelay:
		return "buffer delay"
	case binary.ModuleIDAllPassFilter:
		return "all-pass filter"
	case binary.ModuleIDLooper:
		return "looper"
	case binary.ModuleIDAudioInSwitch:
		return "audio in switch"
	case binary.ModuleIDAudioOutSwitch:
		return "audio out switch"
	case binary.ModuleIDMultiNoise:
		return "noise"
	case binary.ModuleIDStereoSpread:
		return "stereo spread"
	case binary.ModuleIDMultiAudioPanner:
		return "audio panner"
	case binary.ModuleIDPitchShifter:
		return "pitch shifter"
	case binary.ModuleIDBitModulator:
		return "bit modulator"
	case binary.ModuleIDAudioBalance:
		return "audio balance"
	case binary.ModuleIDInverter:
		return "inverter"
	case binary.ModuleIDAudioMixer:
		return "audio mixer"
	case binary.ModuleIDDiffuser:
		return "diffuser"
	case binary.ModuleIDGranular:
		return "granular"
	// Control modules
	case binary.ModuleIDSequencer:
		return "sequencer"
	case binary.ModuleIDLFO:
		return "lfo"
	case binary.ModuleIDASDR:
		return "adsr"
	case binary.ModuleIDSampleAndHold:
		return "sample and hold"
	case binary.ModuleIDCVInvert:
		return "cv invert"
	case binary.ModuleIDSteps:
		return "steps"
	case binary.ModuleIDSlewLimiter:
		return "slew limiter"
	case binary.ModuleIDMultiplier:
		return "multiplier"
	case binary.ModuleIDQuantizer:
		return "quantizer"
	case binary.ModuleIDInSwitch:
		return "in switch"
	case binary.ModuleIDOutSwitch:
		return "out switch"
	case binary.ModuleIDRhythm:
		return "rhythm"
	case binary.ModuleIDRandom:
		return "random"
	case binary.ModuleIDValue:
		return "value"
	case binary.ModuleIDCVDelay:
		return "cv delay"
	case binary.ModuleIDLoop:
		return "loop"
	case binary.ModuleIDCVFilter:
		return "cv filter"
	case binary.ModuleIDClockDivider:
		return "clock divider"
	case binary.ModuleIDComparator:
		return "comparator"
	case binary.ModuleIDCVRectify:
		return "cv rectify"
	case binary.ModuleIDTrigger:
		return "trigger"
	case binary.ModuleIDCVFlipFlop:
		return "cv flip flop"
	// Analysis modules
	case binary.ModuleIDEnvFollower:
		return "env follower"
	case binary.ModuleIDOnsetDetector:
		return "onset detector"
	case binary.ModuleIDPitchDetector:
		return "pitch detector"
	// Effect modules
	case binary.ModuleIDODandDistorsion:
		return "od & distorsion"
	case binary.ModuleIDCompressor:
		return "compressor"
	case binary.ModuleIDPlateReverb:
		return "plate reverb"
	case binary.ModuleIDPhaser:
		return "phaser"
	case binary.ModuleIDGate:
		return "gate"
	case binary.ModuleIDTremolo:
		return "tremolo"
	case binary.ModuleIDToneControl:
		return "tone control"
	case binary.ModuleIDDelayWMod:
		return "delay w/mod"
	case binary.ModuleIDFuzz:
		return "fuzz"
	case binary.ModuleIDGhostverb:
		return "ghostverb"
	case binary.ModuleIDCabinetSim:
		return "cabinet sim"
	case binary.ModuleIDFlanger:
		return "flanger"
	case binary.ModuleIDChorus:
		return "chorus"
	case binary.ModuleIDVibrato:
		return "vibrato"
	case binary.ModuleIDEnvFilter:
		return "env filter"
	case binary.ModuleIDRingModulator:
		return "ring modulator"
	case binary.ModuleIDHallReverb:
		return "hall reverb"
	case binary.ModuleIDPingpongDelay:
		return "pingpong delay"
	case binary.ModuleIDReverbLite:
		return "reverb lite"
	case binary.ModuleIDRoomReverb:
		return "room reverb"
	default:
		return strconv.Itoa(int(t))
	}
}

// MarshalYAML implements yaml.Marshaler.
func (t Type) MarshalYAML() (interface{}, error) {
	return t.String(), nil
}
