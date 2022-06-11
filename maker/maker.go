package maker

import (
	"strings"

	"github.com/Gealber/cvmaker/layout"
	"github.com/jung-kurt/gofpdf"
)

// Maker handle the creation of cv.
type Maker struct {
	pdf *gofpdf.Fpdf
	cv  *CV
	ly  *layout.Layout
}

// NewCVMaker ...
func NewCVMaker() (*Maker, error) {
	cv := newCV()
	pdf := gofpdf.New("P", "mm", "A4", "")

	ly, err := layout.DefaultLayout()
	if err != nil {
		return nil, err
	}

	return &Maker{
		pdf: pdf,
		cv:  cv,
		ly:  ly,
	}, nil
}

// Generate new cv, with data declared in resume.json file.
// Consult documentation for the format of json.
func (m *Maker) Generate() error {
	// title.
	m.pdf.SetTitle(m.cv.Basics.Name, false)
	// author.
	m.pdf.SetAuthor(m.cv.Basics.Name, false)

	m.pdf.AddPage()

	if err := m.paintLayout(); err != nil {
		return err
	}

	return m.pdf.OutputFileAndClose("resume.pdf")
}

func (m *Maker) paintLayout() error {
	for _, region := range m.ly.Regions {
		switch region.Name {
		case "Header":
			m.setHeader(region)
		case "Image":
			m.setImageProfile(region)
		case "Summary":
			m.setSummary(region)
		case "Social Networks":
			m.setNetworks(region)
		case "Education":
			m.setEducation(region)
		case "Work":
			m.setExperience(region)
		}

		if m.pdf.Error() != nil {
			return m.pdf.Error()
		}
	}

	m.setProjects()
	m.setDefaultDrawColor()

	return nil
}

func (m *Maker) setXYFromRegion(region layout.Region) (float64, float64) {
	x := region.TopLeft.X
	y := region.TopLeft.Y
	m.pdf.SetXY(x, y)

	return x, y
}

func (m *Maker) setHeader(region layout.Region) {
	x, y := m.setXYFromRegion(region)

	m.pdf.SetFont("Times", "B", defaultHeaderFontSize)
	m.pdf.Cell(x, y, m.cv.Basics.Name)

	// seting label.
	y += defaultStep
	m.pdf.SetY(y)
	m.setColor(redSoft)
	m.pdf.SetFont("Times", "", largeFontSize)
	m.pdf.Cell(x, y, m.cv.Basics.Label)

	y += defaultStep
	m.pdf.SetY(y)
	m.setColor(grey)
	m.pdf.Cell(x, y, m.cv.Basics.Years+" years of experience")

	// re-setting font.
	m.setDefaultTextColor()
	m.pdf.SetFont("Times", "", defaultFontSize)
}

func (m *Maker) setSummary(region layout.Region) {
	m.setXYFromRegion(region)

	for _, description := range m.cv.Basics.Summary {
		bullet := m.pdf.UnicodeTranslatorFromDescriptor("")("•") + " "
		line := bullet + cleanLine(description)
		m.pdf.MultiCell(region.Width(), smalltStep, line, "", "L", false)
	}
}

func (m *Maker) setNetworks(region layout.Region) {
	x, y := m.setXYFromRegion(region)
	font, _ := m.pdf.GetFontSize()

	m.setColor(redSoft)
	m.pdf.SetFont("Times", "", mediumFontSize)
	m.pdf.Cell(region.Width(), mediumFontSize, "Social Networks")
	m.setDefaultTextColor()
	m.pdf.SetFont("Times", "", font)

	m.pdf.SetXY(x, y+defaultStep)

	for _, prof := range m.cv.Basics.Profiles {
		m.pdf.Cell(region.Width(), font, prof.Network+": "+prof.URL)
		m.pdf.SetXY(x, m.pdf.GetY()+defaultStep)
	}
}

func (m *Maker) setEducation(region layout.Region) {
	x, y := m.setXYFromRegion(region)
	font, _ := m.pdf.GetFontSize()

	m.setColor(redSoft)
	m.pdf.SetFont("Times", "", mediumFontSize)
	m.pdf.Cell(region.Width(), mediumFontSize, "Education")
	m.setDefaultTextColor()
	m.pdf.SetFont("Times", "", font)

	y += defaultStep
	m.pdf.SetXY(x, y)

	for _, edu := range m.cv.Education {
		txt := strings.Join([]string{edu.Institution, edu.StudyType, edu.Area}, ", ")
		m.pdf.Cell(region.Width(), font, txt)

		y += defaultStep
		m.pdf.SetXY(x, y)
	}
}

func (m *Maker) setExperience(region layout.Region) {
	x, y := m.setXYFromRegion(region)
	font, _ := m.pdf.GetFontSize()

	m.setColor(redSoft)
	m.pdf.SetFont("Times", "", mediumFontSize)
	m.pdf.Cell(region.Width(), mediumFontSize, "Work Experience")
	m.setDefaultTextColor()
	m.pdf.SetFont("Times", "", font)

	y += defaultStep
	m.pdf.SetXY(x, y)

	for _, work := range m.cv.Work {
		m.setWork(work)
	}
}

func (m *Maker) setWork(work Work) {
	pageWidth, _ := m.pdf.GetPageSize()

	m.pdf.SetFont("Times", "B", defaultFontSize)
	m.pdf.Cell(defaultStep, defaultStep, work.Name)
	m.pdf.SetFont("Times", "", defaultFontSize)
	// location.
	m.pdf.SetY(m.pdf.GetY() + defaultStep)
	m.pdf.Cell(defaultStep, defaultStep, "Location: "+work.Location)
	// position.
	m.pdf.SetY(m.pdf.GetY() + defaultStep)
	m.pdf.Cell(defaultStep, defaultStep, "Position: "+work.Position)
	m.pdf.SetY(m.pdf.GetY() + defaultStep)
	// highlights.
	m.setBulletPoints(work.HighLights)
	// duration.
	m.pdf.Cell(defaultStep, defaultStep, "Start Date: "+work.StartDate)
	m.pdf.SetY(m.pdf.GetY() + defaultStep)
	m.pdf.Cell(defaultStep, defaultStep, "End Date: "+work.EndDate)
	// website.
	if work.URL != "" {
		m.pdf.SetY(m.pdf.GetY() + defaultStep)
		m.pdf.Cell(defaultStep, defaultStep, "Company Website: "+work.URL)
	}
	// tech stack.
	if len(work.Stack) > 0 {
		stack := strings.Join(work.Stack, ", ")

		m.pdf.SetY(m.pdf.GetY() + defaultStep)
		m.pdf.MultiCell(
			pageWidth-3*defaultMargin,
			smallFontSize,
			"Tech Stack: "+stack,
			"", "L", false)
	}

	m.pdf.SetY(m.pdf.GetY() + defaultStep)
	m.pdf.SetLineWidth(defaultLineWidth)
	m.setDrawColor(redSoft)
	m.pdf.Line(
		m.pdf.GetX(),
		m.pdf.GetY(),
		pageWidth-defaultMargin,
		m.pdf.GetY())
	m.pdf.SetX(defaultMargin)
}

func (m *Maker) setProjects() {
	m.setMediumLabel("Personal Projects")
	m.pdf.SetY(m.pdf.GetY() + defaultStep)
	pageWidth, _ := m.pdf.GetPageSize()

	for _, project := range m.cv.Projects {
		m.pdf.SetFont("Times", "B", defaultFontSize)
		m.pdf.Cell(defaultStep, defaultStep, project.Name)
		m.pdf.SetFont("Times", "", defaultFontSize)
		// highlights.
		m.pdf.SetY(m.pdf.GetY() + defaultStep)
		m.setBulletPoints(project.HighLights)
		// website.
		if project.URL != "" {
			m.pdf.Cell(defaultStep, defaultStep, "Project website: "+project.URL)
		}

		m.pdf.SetY(m.pdf.GetY() + defaultStep)
		m.pdf.SetLineWidth(defaultLineWidth)
		m.setDrawColor(redSoft)
		m.pdf.Line(
			m.pdf.GetX(),
			m.pdf.GetY(),
			pageWidth-defaultMargin,
			m.pdf.GetY())
		m.pdf.SetX(defaultMargin)
	}
}

func (m *Maker) setImageProfile(region layout.Region) {
	x, y := m.setXYFromRegion(region)

	options := gofpdf.ImageOptions{
		ReadDpi:   false,
		ImageType: "JPG",
	}

	m.pdf.ImageOptions(
		"profile.jpg",
		x,
		y,
		region.Width(), region.Height(),
		false,
		options,
		0,
		m.cv.Basics.Image,
	)
}

func (m *Maker) setBulletPoints(points []string) {
	pageWidth, _ := m.pdf.GetPageSize()

	for _, description := range points {
		bullet := m.pdf.UnicodeTranslatorFromDescriptor("")("•") + " "
		line := bullet + cleanLine(description)
		m.pdf.MultiCell(pageWidth-3*defaultMargin, smalltStep, line, "", "L", false)
	}
}

func (m *Maker) setMediumLabel(txt string) {
	font, _ := m.pdf.GetFontSize()
	m.pdf.SetFont("Times", "", mediumFontSize)
	m.setColor(redSoft)
	m.pdf.Cell(defaultCellWidth, defaultMargin, txt)
	m.setDefaultTextColor()
	m.pdf.SetFont("Times", "", font)
}

func (m *Maker) setDefaultTextColor() {
	m.setColor(black)
}

func (m *Maker) setColor(c colorRGB) {
	m.pdf.SetTextColor(c[0], c[1], c[2])
}

func (m *Maker) setDrawColor(c colorRGB) {
	m.pdf.SetDrawColor(c[0], c[1], c[2])
}

func (m *Maker) setDefaultDrawColor() {
	m.setDrawColor(black)
}
