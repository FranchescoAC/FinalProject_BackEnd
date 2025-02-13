namespace UpdatePayment.model
{
    public class Payment
    {
        public int Id { get; set; }
        public decimal Amount { get; set; }
        public string Description { get; set; }
        public DateTime Date { get; set; }

        public Payment()
        {
            Description = string.Empty;
            Date = DateTime.Now; 
        }
    }
}
