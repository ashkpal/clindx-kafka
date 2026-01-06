package producer

func (p *Producer) Close() error {
	return p.writer.Close()
}
