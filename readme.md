# WeatherDump

Multi-platform software for record, demodulate, decode and process data from weather satellites.

## Supported Datalink Protocols

| Protocol | Complete Name | Satellites | Band | Support Level |
| -------- | ------------- | ---------- | ---- | ------------- |
| LRPT | Low Rate Picture Transfer | Meteor-MN2 | VHF | Alpha |
| HRD | High Rate Data | NOAA-20 & Suomi | X-Band | Beta |
| APT | Automatic Picture Transfer | NOAA-15, NOAA-18 & NOAA-19 | VHF | Planned (Beta 1) |

## Example Usage

Decoding and processing a Meteor-MN2 soft-symbol file:

```bash
weatherdump lrpt soft ./file_path.bin
```

## Upcoming Features List

- [ ] Support SatNOGS compatible output.
- [ ] Add multi-thread support to decoder.
- [ ] Generate LRPT RGB composite.
- [ ] Add NOAA APT support.

## Changelog

### Alpha 2

- Support for multiple decoders for each protocol.
- Added support to scrambled unsynchronized CADU files.
- Added support to synchronized unscrambled CADU files.
- Fix multi-thread image processing freezing.
- New decoder and processor progress indicator.
- New CLI argument for multiple decoders.
- Improved far from perfect documentation.
- First public release of GUI version.
- Statistical SCID recover.