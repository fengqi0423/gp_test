package gp

type RawSample struct {
    Features map[string]string
    Label int

    Prediction float64
}

func NewRawSample() *RawSample {
    ret := RawSample{}
    ret.Features = make(map[string]string)
    ret.Label = 0
    ret.Prediction = 0.0
    return &ret
}

func (s *RawSample) GetFeatureValue(key string) string {
    value, ok := s.Features[key]
    if ok {
        return value
    } else {
        return "nil"
    }
}


/*
Here, label should be int value started from 0
*/
type Sample struct {
    Features []Feature
    Label int

    Prediction float64
}

func NewSample() *Sample {
    ret := Sample{}
    ret.Features = []Feature{}
    ret.Label = 0
    ret.Prediction = 0.0
    return &ret
}

func (s *Sample) Clone() *Sample {
    ret := NewSample()
    ret.Label = s.Label
    ret.Prediction = s.Prediction
    for _, feature := range s.Features {
        clone_feature := Feature{feature.Id, feature.Value}
        ret.Features = append(ret.Features, clone_feature)
    }

    return ret
}
