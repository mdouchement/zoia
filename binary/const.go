package binary

const (
	// SizeHeader is the length of the patch's header.
	SizeHeader = 24
	// SizeCaption is the length of a name.
	SizeCaption = 15
	// SizeControl is the size of a module control.
	SizeControl = 4
	// SizeNumberOfPatchCables is the length for storing the number of patch cables.
	SizeNumberOfPatchCables = 4
	// SizePatchCable is the length of a patch cable between 2 modules.
	SizePatchCable = 20
	// SizeNumberOfPages is the length for storing the number of pages details.
	SizeNumberOfPages = 4
	// SizePage is the length of the page details.
	SizePage = 16
	// SizePagePadding is the length of the padding at the end of pages data.
	SizePagePadding = 4
	// SizeModuleExtra is the length of extra modules data that are at the end of the patch.
	SizeModuleExtra = 4
)

// All these offsets are relative to their block (header, module, etc.).
const (
	// OffsetHeaderName is the header name.
	OffsetHeaderName = 4
	// OffsetHeaderNumberOfModules is the number of modules in the patch.
	OffsetHeaderNumberOfModules = 20

	// OffsetModuleID is the module ID (vca, lfo, etc.).
	OffsetModuleID = 4
	// OffsetModulePage is the page on which the module is.
	OffsetModulePage = 12
	// OffsetModuleColor is the color of the module (seems to be a legacy value).
	OffsetModuleColor = 16
	// OffsetModulePosition is the starting position of the module on the grid.
	OffsetModulePosition = 20
	// OffsetModuleChannel is the used channel (stereo, left and right).
	OffsetModuleChannel = 33

	// OffsetPatchCableModuleSource is the index of the source module in the module list.
	OffsetPatchCableModuleSource = 0
	// OffsetPatchCableSignalSource is the source type of the intput patch cable (left, right, CV, etc.).
	OffsetPatchCableSignalSource = 4
	// OffsetPatchCableModuleSource is the index of the destination module in the module list.
	OffsetPatchCableModuleDestination = 8
	// OffsetPatchCableSignalDestination is the destination type of the intput patch cable (left, right, CV, etc.).
	OffsetPatchCableSignalDestination = 12
	// OffsetPatchCableGain is the gain of the patch cable.
	OffsetPatchCableGain = 16

	// OffsetModuleExtraColor is the color of the module (implemeted in version 1.10)
	OffsetModuleExtraColor = 0

	// OffsetTailNumberOfPatchCables is the number of patch cables.
	OffsetTailNumberOfPatchCables = 0
)

// Dedicated module offsets
const (
	// OffsetAudioInputChannel is the used channel (stereo, left and right).
	OffsetAudioInputChannel = 32

	// OffsetAudioOutputChannel is the used channel (stereo, left and right).
	OffsetAudioOutputChannel = 33
	// OffsetAudioOutputGain is the gain of the module.
	OffsetAudioOutputGain = 40

	// OffsetVCASignal is signal configuration (mono/stereo).
	OffsetVCASignal = 32
	// OffsetVCALevelControl is the amplification of the VCA.
	OffsetVCALevelControl = 40

	// OffsetCompressorNbOfControls is number of activated controls (treshold is always enabled).
	OffsetCompressorNbOfControls = 24
	// OffsetCompressorAttackEnabled informs when attack control is enabled.
	OffsetCompressorAttackEnabled = 32
	// OffsetCompressorReleaseEnabled informs when release control is enabled.
	OffsetCompressorReleaseEnabled = 33
	// OffsetCompressorRatioEnabled informs when ratio control is enabled.
	OffsetCompressorRatioEnabled = 34
	// OffsetCompressorSignal is signal configuration (mono/stereo).
	OffsetCompressorSignal = 35
	// OffsetCompressorSidechain is sidechain configuration (internal/external).
	OffsetCompressorSidechain = 36
	// OffsetCompressorControlsStart is where the control values start.
	OffsetCompressorControlsStart = 40
)

// Identifiers off all the modules.
const (
	// Interface modules
	ModuleIDAudioInput    = 1
	ModuleIDAudioOutput   = 2
	ModuleIDPushbutton    = 15
	ModuleIDKeyboard      = 16
	ModuleIDMidiNotesIn   = 20
	ModuleIDMidiCCIn      = 21
	ModuleIDMidiPressure  = 35
	ModuleIDStompswitch   = 44
	ModuleIDCportExpCVIn  = 54
	ModuleIDCportExpCVOut = 55
	ModuleIDUIButton      = 56
	ModuleIDMidiNoteOut   = 60
	ModuleIDMidiCCOut     = 61
	ModuleIDMidiPCOut     = 62
	ModuleIDPixel         = 81
	ModuleIDMidiClockIn   = 82
	// Audio modules
	ModuleIDSVFilter         = 0
	ModuleIDSVAliaser        = 3
	ModuleIDVCA              = 7
	ModuleIDAudionMultiply   = 8
	ModuleIDBitCrusher       = 9
	ModuleIDDelayLine        = 13
	ModuleIDOscillator       = 14
	ModuleIDMultiFilter      = 24
	ModuleIDBufferDelay      = 26
	ModuleIDAllPassFilter    = 27
	ModuleIDLooper           = 30
	ModuleIDAudioInSwitch    = 33
	ModuleIDAudioOutSwitch   = 34
	ModuleIDMultiNoise       = 38
	ModuleIDStereoSpread     = 53
	ModuleIDMultiAudioPanner = 57
	ModuleIDPitchShifter     = 59
	ModuleIDBitModulator     = 63
	ModuleIDAudioBalance     = 64
	ModuleIDInverter         = 65
	ModuleIDAudioMixer       = 76
	ModuleIDDiffuser         = 78
	ModuleIDGranular         = 83
	// Control modules
	ModuleIDSequencer     = 4
	ModuleIDLFO           = 5
	ModuleIDASDR          = 6
	ModuleIDSampleAndHold = 10
	ModuleIDCVInvert      = 17
	ModuleIDSteps         = 18
	ModuleIDSlewLimiter   = 19
	ModuleIDMultiplier    = 22
	ModuleIDQuantizer     = 28
	ModuleIDInSwitch      = 31
	ModuleIDOutSwitch     = 32
	ModuleIDRhythm        = 37
	ModuleIDRandom        = 39
	ModuleIDValue         = 45
	ModuleIDCVDelay       = 46
	ModuleIDLoop          = 47
	ModuleIDCVFilter      = 48
	ModuleIDClockDivider  = 49
	ModuleIDComparator    = 50
	ModuleIDCVRectify     = 51
	ModuleIDTrigger       = 52
	ModuleIDCVFlipFlop    = 77
	// Analysis modules
	ModuleIDEnvFollower   = 12
	ModuleIDOnsetDetector = 36
	ModuleIDPitchDetector = 58
	// Effect modules
	ModuleIDODandDistorsion = 11
	ModuleIDCompressor      = 23
	ModuleIDPlateReverb     = 25
	ModuleIDPhaser          = 29
	ModuleIDGate            = 40
	ModuleIDTremolo         = 41
	ModuleIDToneControl     = 42
	ModuleIDDelayWMod       = 43
	ModuleIDFuzz            = 66
	ModuleIDGhostverb       = 67
	ModuleIDCabinetSim      = 68
	ModuleIDFlanger         = 69
	ModuleIDChorus          = 70
	ModuleIDVibrato         = 71
	ModuleIDEnvFilter       = 72
	ModuleIDRingModulator   = 73
	ModuleIDHallReverb      = 74
	ModuleIDPingpongDelay   = 75
	ModuleIDReverbLite      = 79
	ModuleIDRoomReverb      = 80
)

// Channels / Signals.
const (
	ChannelStereo = 0
	ChannelLeft   = 1
	ChannelRight  = 2

	SignalMono   = 0
	SignalStereo = 1
)

// Uncategorized values.
const (
	SidechainInternal = 0
	SidechainExternal = 1
)

// Colors
const (
	// From module info (legacy)
	// In the cycle order defined by Zoia
	ColorRed     = 3
	ColorOrange  = 3
	ColorMango   = 4
	ColorYello   = 4
	ColorLime    = 2
	ColorGreen   = 2
	ColorSurf    = 5
	ColorAqua    = 5
	ColorSky     = 1
	ColorBlue    = 1
	ColorPurple  = 6
	ColorMagenta = 6
	ColorPink    = 3
	ColorPeach   = 3
	ColorWhite   = 7

	// From extra module data
	ColorExtraBlue    = 1
	ColorExtraGreen   = 2
	ColorExtraRed     = 3
	ColorExtraYello   = 4
	ColorExtraAqua    = 5
	ColorExtraMagenta = 6
	ColorExtraWhite   = 7
	ColorExtraOrange  = 8
	ColorExtraLime    = 9
	ColorExtraSky     = 11
	ColorExtraPurple  = 12
	ColorExtraSurf    = 10
	ColorExtraPink    = 13
	ColorExtraPeach   = 14
	ColorExtraMango   = 15
)
