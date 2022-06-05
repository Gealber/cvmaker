package maker

import (
	"strings"

	"github.com/jung-kurt/gofpdf"
)

// Maker handle the creation of cv.
type Maker struct {
	pdf *gofpdf.Fpdf
	cv  *CV
}

// NewCVMaker ...
func NewCVMaker() *Maker {
	cv := newCV()
	pdf := gofpdf.New("P", "mm", "A4", "")

	return &Maker{pdf: pdf, cv: cv}
}

// Generate new cv, with data declared in resume.json file.
// Consult documentation for the format of json.
func (m *Maker) Generate() error {
	// title.
	m.pdf.SetTitle(m.cv.Basics.Name, false)
	// author.
	m.pdf.SetAuthor(m.cv.Basics.Name, false)

	m.pdf.AddPage()

	m.pdf.SetLeftMargin(defaultMargin)
	m.pdf.SetRightMargin(defaultMargin)
	m.pdf.SetTopMargin(defaultMargin)

	// setting some top margin.
	m.pdf.SetXY(defaultMargin, defaultMargin)

	// setting name.
	m.setName()

	// seting label.
	m.pdf.SetY(2*defaultMargin + smalltStep)
	m.setColor(redSoft)
	m.setLabel()
	m.pdf.SetY(m.pdf.GetY() + defaultStep)
	m.setColor(grey)
	m.pdf.Cell(defaultCellWidth, defaultMargin, m.cv.Basics.Years+" years of experience")

	// setting image.
	m.setImageProfile()

	// setting font.
	m.setDefaultTextColor()
	m.pdf.SetFont("Times", "", defaultFontSize)

	// setting summary.
	m.pdf.SetXY(defaultMargin, smalltStep*defaultMargin)
	m.setBulletPoints(m.cv.Basics.Summary)

	// setting networks.
	m.pdf.SetY(m.pdf.GetY() + defaultStep)
	m.setNetworks()

	// setting education.
	m.pdf.SetY(m.pdf.GetY() + defaultStep + smalltStep)
	m.setEducation()

	// setting work experience.
	m.pdf.SetY(m.pdf.GetY() + defaultStep)
	m.pdf.SetY(m.pdf.GetY() + defaultStep)
	m.setExperience()

	// personal projects.
	m.pdf.SetY(m.pdf.GetY() + defaultStep)
	m.setProjects()

	return m.pdf.OutputFileAndClose("resume.pdf")
}

func (m *Maker) setNetworks() {
	m.setMediumLabel("Social Networks")

	_, height := m.pdf.GetXY()
	m.pdf.SetXY(defaultMargin, height-smalltStep)

	for _, prof := range m.cv.Basics.Profiles {
		m.pdf.Cell(defaultCellWidth, defaultCellHeight, prof.Network+": "+prof.URL)
		m.pdf.SetXY(defaultMargin, m.pdf.GetY()+defaultStep)
	}
}

func (m *Maker) setExperience() {
	m.setMediumLabel("Work Experience")
	m.pdf.SetY(m.pdf.GetY() + defaultStep)
	pageWidth, _ := m.pdf.GetPageSize()

	for _, work := range m.cv.Work {
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
			m.pdf.GetX()+pageWidth-3*defaultMargin, m.pdf.GetY())
		m.pdf.SetX(defaultMargin)
	}
}

func (m *Maker) setEducation() {
	m.setMediumLabel("Education")

	_, height := m.pdf.GetXY()
	m.pdf.SetXY(defaultMargin, height-smalltStep)

	for _, edu := range m.cv.Education {
		txt := strings.Join([]string{edu.Institution, edu.StudyType, edu.Area}, ", ")
		m.pdf.Cell(defaultCellWidth, defaultCellHeight, txt)
		m.pdf.SetXY(defaultMargin, m.pdf.GetY()+defaultStep)
	}
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
			m.pdf.GetX()+pageWidth-3*defaultMargin, m.pdf.GetY())
		m.pdf.SetX(defaultMargin)
	}
}

func (m *Maker) setLabel() {
	m.pdf.SetFont("Times", "", largeFontSize)
	m.pdf.Cell(defaultCellWidth, defaultMargin, m.cv.Basics.Label)
}

func (m *Maker) setName() {
	m.pdf.SetFont("Times", "B", defaultImageSize)
	m.pdf.Cell(defaultCellWidth, defaultMargin, m.cv.Basics.Name)
}

func (m *Maker) setImageProfile() {
	pageWidth, _ := m.pdf.GetPageSize()
	options := gofpdf.ImageOptions{
		ReadDpi:   false,
		ImageType: "JPG",
	}

	m.pdf.ImageOptions(
		"profile.jpg",
		pageWidth-defaultMargin-defaultImageSize,
		defaultMargin,
		defaultImageSize, defaultImageSize,
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
